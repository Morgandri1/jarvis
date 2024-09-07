import Axios, { AxiosRequestConfig } from "axios";
import { Result } from "./types/api";
import { AuthResponse } from "./types/auth";

const axios_config: AxiosRequestConfig = {
  baseURL: "http://localhost:80",
  method: "POST",
  headers: {
    "Content-Type": "application/json",
    client: "JARVIS-MOBILE"
  }
}

export default class User {
  public id: string;
  public username: string;
  private password: string;

  private constructor(id: string, username: string, password: string) {
    this.id = id;
    this.username = username;
    this.password = password;
  }

  static async login(username: string, password: string): Promise<User> {
    const response = await Axios({
      ...axios_config,
      url: "/login",
      data: { username, password }
    });

    const { data }: Result<AuthResponse> = response.data

    return new User(data.user.id, data.user.username, password);
  }
}
