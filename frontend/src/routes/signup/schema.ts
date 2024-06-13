import { z } from 'zod';

export const signupSchema = z.object({
  email: z.string().email(),
  password: z.string(),
});

export type SignupSchema = typeof signupSchema;

export type SignupForm = z.infer<typeof signupSchema>;
