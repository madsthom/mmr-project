import { z } from 'zod';

const teamSchema = z.object({
  // TODO: Get all values from server or configuration file
  score: z.number().int().positive().max(10),
  member1: z.string().max(4).toLowerCase(),
  member2: z.string().max(4).toLowerCase(),
});

export const matchSchema = z.object({
  team1: teamSchema,
  team2: teamSchema,
});

export type MatchSchema = typeof matchSchema;

export type MatchForm = z.infer<typeof matchSchema>;
