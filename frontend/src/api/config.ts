import axios, {
  type AxiosRequestConfig,
  type RawAxiosRequestHeaders,
} from "axios";

const port = import.meta.env.VITE_PORT ?? 3000;
const API = import.meta.env.VITE_API ?? "localhost";

const token = localStorage.getItem("token");

export const client = axios.create({
  baseURL: `${API}:${port}/api`,
});

export const config: AxiosRequestConfig = {
  headers: {
    Accept: "application/json",
    Authorization: token,
  } as RawAxiosRequestHeaders,
};
