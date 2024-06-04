import { z } from 'zod';

const teamSchema = z.object({
  // TODO: Get all values from server or configuration file
  score: z.number().int().min(0).max(10),
  member1: z.number().int(),
  member2: z.number().int(),
});

export const matchSchema = z.object({
  team1: teamSchema,
  team2: teamSchema,
});

export type MatchSchema = typeof matchSchema;

export type MatchForm = z.infer<typeof matchSchema>;
