using Microsoft.AspNetCore.Mvc;
using MMRProject.Api.DTOs;

namespace MMRProject.Api.Controllers;

[ApiController]
[Route("api/v2/mmr")]
public class MMRV2Controller(ILogger<MMRV2Controller> logger) : ControllerBase
{
    [HttpGet("matches")]
    public async Task<IEnumerable<MatchDetailsV2>> GetMatches([FromQuery] int limit = 100, [FromQuery] int offset = 0)
    {
        return
        [
            new MatchDetailsV2
            {
                Date = DateTimeOffset.Now,
                Team1 = new MatchTeamV2()
                {
                    Score = 1,
                    Member1 = "Player1",
                    Member2 = "Player2"
                },
                Team2 = new MatchTeamV2()
                {
                    Score = 0,
                    Member1 = "Player3",
                    Member2 = "Player4"
                },
                MMRCalculation = new MatchMMRCalculationDetails
                {
                    Team1 = new MatchMMRCalculationTeam()
                    {
                        Player1MMRDelta = 10,
                        Player2MMRDelta = -10
                    },
                    Team2 = new MatchMMRCalculationTeam()
                    {
                        Player1MMRDelta = -10,
                        Player2MMRDelta = 10
                    }
                }
            }
        ];
    }
    
    [HttpPost("matches")]
    public async Task SubmitMatch()
    {

    }
}