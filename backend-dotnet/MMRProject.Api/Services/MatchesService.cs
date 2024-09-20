using Microsoft.EntityFrameworkCore;
using MMRProject.Api.Data;
using MMRProject.Api.Data.Entities;
using MMRProject.Api.DTOs;

namespace MMRProject.Api.Services;

public interface IMatchesService
{
    Task<IEnumerable<Match>> GetMatchesForSeason(long seasonId, int limit, int offset, bool orderByCreatedAtDescending,
        bool includeMmrCalculations, long? userId);

    Task SubmitMatch(long seasonId, SubmitMatchV2Request request);

    Task RecalculateMMRForMatchesInSeason(long seasonId, long? fromMatchId = null);
}

public class MatchesService(ApiDbContext dbContext, IUserService userService) : IMatchesService
{
    public async Task<IEnumerable<Match>> GetMatchesForSeason(long seasonId, int limit, int offset,
        bool orderByCreatedAtDescending,
        bool includeMmrCalculations, long? userId)
    {
        var query = dbContext.Matches
            .Include(x => x.TeamOne)
            .Include(x => x.TeamTwo)
            .Where(x => x.SeasonId == seasonId);

        if (includeMmrCalculations)
        {
            query = query.Include(x => x.MmrCalculations);
        }

        if (orderByCreatedAtDescending)
        {
            query = query.OrderByDescending(x => x.CreatedAt);
        }
        else
        {
            query = query.OrderBy(x => x.CreatedAt);
        }

        if (userId.HasValue)
        {
            query = query.Where(x => x.TeamOne!.UserOneId == userId || x.TeamOne.UserTwoId == userId ||
                                     x.TeamTwo!.UserOneId == userId || x.TeamTwo.UserTwoId == userId);
        }

        return await query.Skip(offset).Take(limit).ToListAsync();
    }

    public async Task SubmitMatch(long seasonId, SubmitMatchV2Request request)
    {
        var uniquePlayers = new HashSet<long>
        {
            request.Team1.Member1,
            request.Team1.Member2,
            request.Team2.Member1,
            request.Team2.Member2
        };

        if (uniquePlayers.Count != 4)
        {
            // TODO: Throw better exception
            throw new Exception("Players must be unique");
        }

        var exists = await CheckExistingMatch(request.Team1.Member1, request.Team1.Member2, request.Team2.Member1,
            request.Team2.Member2, request.Team1.Score, request.Team2.Score);
        if (exists)
        {
            // TODO: Throw better exception
            throw new Exception("Match already exists");
        }

        var players = await dbContext.Users.Where(x => uniquePlayers.Contains(x.Id)).ToListAsync();

        if (players.Count != uniquePlayers.Count)
        {
            // TODO: Throw better exception
            throw new Exception("Not all players were found");
        }

        var match = await CreateMatch(seasonId, request.Team1.Member1, request.Team1.Member2, request.Team2.Member1,
            request.Team2.Member2, request.Team1.Score, request.Team2.Score);

        await CalculateMMR(seasonId, match);
    }

    private async Task CalculateMMR(long seasonId, Match match)
    {
        // TODO: Implement MMR calculation by calling service?
        await Task.Delay(300);
        var mmrCalculationRequest = new MMRCalculationRequest
        {
            Team1 = new MMRCalculationTeam
            {
                Score = (int)match.TeamOne!.Score!, // TODO: Fix this
                Players =
                [
                    await PlayerRatingForUserAsync(match.TeamOne!.UserOneId!.Value, seasonId),
                    await PlayerRatingForUserAsync(match.TeamOne!.UserTwoId!.Value, seasonId)
                ]
            },
            Team2 = new MMRCalculationTeam
            {
                Score = (int)match.TeamTwo!.Score!,
                Players =
                [
                    await PlayerRatingForUserAsync(match.TeamTwo!.UserOneId!.Value, seasonId),
                    await PlayerRatingForUserAsync(match.TeamTwo!.UserTwoId!.Value, seasonId)
                ]
            }
        };

        // TODO: Call API to calculate MMR

        // TODO: Insert player histories
        // TODO: Insert MMR calculations
    }

    private async Task<MMRCalculationPlayerRating> PlayerRatingForUserAsync(long userId, long seasonId)
    {
        var playerHistory = await userService.LatestPlayerHistoryAsync(userId, seasonId);

        return new MMRCalculationPlayerRating
        {
            Id = userId,
            Mu = playerHistory?.Mu,
            Sigma = playerHistory?.Sigma
        };
    }

    private async Task<bool> CheckExistingMatch(long playerOneId, long playerTwoId, long playerThreeId,
        long playerFourId, int teamOneScore, int teamTwoScore)
    {
        // TODO: Validate this is the correct logic
        var exists = await dbContext.Matches
            .AsNoTracking()
            .Where(m =>
                (
                    (m.TeamOne!.UserOneId == playerOneId || m.TeamOne.UserTwoId == playerOneId) &&
                    (m.TeamOne.UserOneId == playerTwoId || m.TeamOne.UserTwoId == playerTwoId) &&
                    m.TeamOne.Score == teamOneScore &&
                    (m.TeamTwo!.UserOneId == playerThreeId || m.TeamTwo.UserTwoId == playerThreeId) &&
                    (m.TeamTwo.UserOneId == playerFourId || m.TeamTwo.UserTwoId == playerFourId) &&
                    m.TeamTwo.Score == teamTwoScore
                ) ||
                (
                    (m.TeamOne.UserOneId == playerThreeId || m.TeamOne.UserTwoId == playerThreeId) &&
                    (m.TeamOne.UserOneId == playerFourId || m.TeamOne.UserTwoId == playerFourId) &&
                    m.TeamOne.Score == teamTwoScore &&
                    (m.TeamTwo!.UserOneId == playerOneId || m.TeamTwo.UserTwoId == playerOneId) &&
                    (m.TeamTwo.UserOneId == playerTwoId || m.TeamTwo.UserTwoId == playerTwoId) &&
                    m.TeamTwo.Score == teamOneScore
                ) &&
                m.CreatedAt > DateTime.UtcNow.AddMinutes(-10)
            )
            .Select(m => m.Id)
            .AnyAsync();

        return exists;
    }

    private async Task<Match> CreateMatch(long seasonId, long teamOnePlayerOneId, long teamOnePlayerTwoId,
        long teamTwoPlayerOneId,
        long teamTwoPlayerTwoId, int teamOneScore, int teamTwoScore)
    {
        var match = new Match
        {
            SeasonId = seasonId,
            CreatedAt = DateTime.UtcNow,
            UpdatedAt = DateTime.UtcNow,
            TeamOne = await CreateTeam(teamOnePlayerOneId, teamOnePlayerTwoId, teamOneScore,
                teamOneScore > teamTwoScore),
            TeamTwo = await CreateTeam(teamTwoPlayerOneId, teamTwoPlayerTwoId, teamTwoScore,
                teamTwoScore > teamOneScore)
        };

        await dbContext.Matches.AddAsync(match);
        await dbContext.SaveChangesAsync();

        return match;

        async Task<Team> CreateTeam(long userOneId, long userTwoId, int score, bool isWinner)
        {
            var team = new Team
            {
                UserOneId = userOneId,
                UserTwoId = userTwoId,
                Score = score,
                Winner = isWinner,
                CreatedAt = DateTime.UtcNow,
                UpdatedAt = DateTime.UtcNow
            };

            await dbContext.Teams.AddAsync(team);

            return team;
        }
    }

    public async Task RecalculateMMRForMatchesInSeason(long seasonId, long? fromMatchId)
    {
        await ClearMMRCalculations(seasonId, fromMatchId);

        var matchesQuery = dbContext.Matches
            .Include(x => x.TeamOne)
            .Include(x => x.TeamTwo)
            .Where(x => x.SeasonId == seasonId);

        if (fromMatchId.HasValue)
        {
            matchesQuery = matchesQuery.Where(x => x.Id >= fromMatchId.Value);
        }

        var matches = await matchesQuery.ToListAsync();

        foreach (var match in matches)
        {
            await CalculateMMR(seasonId, match);
        }
    }

    private async Task ClearMMRCalculations(long seasonId, long? fromMatchId)
    {
        var playerHistoryQuery = dbContext.PlayerHistories.Where(x => x.Match!.SeasonId == seasonId);
        if (fromMatchId.HasValue)
        {
            playerHistoryQuery = playerHistoryQuery.Where(x => x.MatchId >= fromMatchId.Value);
        }

        var playerHistories = await playerHistoryQuery.ToListAsync();
        foreach (var playerHistory in playerHistories)
        {
            playerHistory.DeletedAt = DateTime.UtcNow;
        }

        var mmrCalculationsQuery = dbContext.MmrCalculations.Where(x => x.Match!.SeasonId == seasonId);
        if (fromMatchId.HasValue)
        {
            mmrCalculationsQuery = mmrCalculationsQuery.Where(x => x.MatchId >= fromMatchId.Value);
        }

        var calculations = await mmrCalculationsQuery.ToListAsync();
        foreach (var mmrCalculation in calculations)
        {
            mmrCalculation.DeletedAt = DateTime.UtcNow;
        }

        await dbContext.SaveChangesAsync();
    }
}