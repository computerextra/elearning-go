import { z } from "zod";

const AuthResponse = z.object({
  auth: z.boolean(),
});
type AuthResponse = z.infer<typeof AuthResponse>;

const LoginResponse = z.object({
  Authorization: z.string(),
});
type LoginResponse = z.infer<typeof LoginResponse>;
