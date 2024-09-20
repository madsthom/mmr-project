namespace MMRProject.Api.DTOs;

public record PlayerHistoryDetails
{
    public required long UserId { get; set; }
    public required string Name { get; set; }
    public DateTimeOffset Date { get; set; }
    public int MMR { get; set; }
}