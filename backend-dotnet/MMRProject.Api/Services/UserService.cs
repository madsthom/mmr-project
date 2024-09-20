using Microsoft.EntityFrameworkCore;
using MMRProject.Api.Data;
using MMRProject.Api.Data.Entities;

namespace MMRProject.Api.Services;

public interface IUserService
{
    Task<List<User>> AllUsersAsync(string? searchQuery = default);
    Task<User> CreateUserAsync(string name, string? displayName);
    Task<User?> GetUserAsync(long userId);
    Task<PlayerHistory?> LatestPlayerHistoryAsync(long userId, long seasonId);
}

public class UserService(ApiDbContext dbContext) : IUserService
{
    public async Task<List<User>> AllUsersAsync(string? searchQuery = default)
    {
        var query = dbContext.Users.AsQueryable();

        if (!string.IsNullOrWhiteSpace(searchQuery))
        {
            var pattern = $"%{searchQuery}%";
            query = query.Where(
                x => EF.Functions.ILike(x.Name!, pattern) || EF.Functions.ILike(x.DisplayName!, pattern));
        }

        return await query.ToListAsync();
    }

    public async Task<User> CreateUserAsync(string name, string? displayName)
    {
        var user = new User
        {
            Name = name,
            DisplayName = displayName,
            CreatedAt = DateTime.UtcNow,
            UpdatedAt = DateTime.UtcNow
        };
        dbContext.Users.Add(user);
        await dbContext.SaveChangesAsync();
        return user;
    }

    public async Task<User?> GetUserAsync(long userId)
    {
        return await dbContext.Users.FindAsync(userId);
    }

    public async Task<PlayerHistory?> LatestPlayerHistoryAsync(long userId, long seasonId)
    {
        return await dbContext.PlayerHistories.Where(x => x.UserId == userId && x.Match!.SeasonId == seasonId)
            .OrderByDescending(x => x.MatchId)
            .FirstOrDefaultAsync();
    }
}