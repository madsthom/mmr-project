import { z } from 'zod';

export const createUserSchema = z.object({
  name: z.string(),
  displayName: z.string().optional(),
});

export type CreateUserSchema = typeof createUserSchema;
