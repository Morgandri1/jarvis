export interface AuthResponse {
  authenticated: boolean,
  token: string,
  user: {
    id: string,
    username: string
  }
}
