namespace MMRProject.Api.Services;

public interface ISeasonService
{
    Task<int> CurrentSeasonIdAsync();
}

public class SeasonService : ISeasonService
{
    public async Task<int> CurrentSeasonIdAsync()
    {
    }
}