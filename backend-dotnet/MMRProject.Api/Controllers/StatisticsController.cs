using Microsoft.AspNetCore.Mvc;
using MMRProject.Api.DTOs;

namespace MMRProject.Api.Controllers;

[ApiController]
[Route("api/v1/stats")]
public class StatisticsController : ControllerBase
{
    [HttpGet("leaderboard")]
    public async Task<IEnumerable<LeaderboardEntry>> GetLeaderboard()
    {
        throw new NotImplementedException();
    }
    
    [HttpGet("player-history")]
    public async Task<IEnumerable<PlayerHistoryDetails>> GetPlayerHistory([FromQuery] long userId)
    {
        throw new NotImplementedException();
    }
    
    [HttpGet("time-distribution")]
    public async Task<IEnumerable<TimeStatisticsEntry>> GetTimeDistribution()
    {
        throw new NotImplementedException();
    }
}