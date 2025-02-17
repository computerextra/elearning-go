import { client } from "./config";
import { AuthResponse } from "./types";

const checkAuth = async (): Promise<AuthResponse> => {
  const res = await client.get<AuthResponse>("/auth");
  return res.data;
};

export { checkAuth };
