using Microsoft.AspNetCore.Mvc;
using MMRProject.Api.DTOs;
using MMRProject.Api.Mappers;
using MMRProject.Api.Services;

namespace MMRProject.Api.Controllers;

[ApiController]
[Route("api/v1/profile")]
public class ProfileController(IUserService userService) : ControllerBase
{
    [HttpGet]
    public async Task<ProfileDetails> GetProfile()
    {
        // TODO: Implement this. Grabbing user from identityUserId
        return new ProfileDetails
        {
            UserId = null
        };
    }

    [HttpPost("claim")]
    public async Task<ProfileDetails> ClaimProfile([FromBody] ClaimProfileRequest request)
    {
        // TODO: Implement this. Claiming profile for user
        return new ProfileDetails
        {
            UserId = request.UserId
        };
    }
}