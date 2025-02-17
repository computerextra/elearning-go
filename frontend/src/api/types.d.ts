import { z } from "zod";

const AuthResponse = z.object({
  auth: z.boolean(),
});
type AuthResponse = z.infer<typeof AuthResponse>;
