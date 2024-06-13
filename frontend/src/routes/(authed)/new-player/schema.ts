import { z } from 'zod';

export const createPlayerSchema = z.object({
  name: z.string(),
  displayName: z.string().optional(),
});

export type CreatePlayerSchema = typeof createPlayerSchema;
