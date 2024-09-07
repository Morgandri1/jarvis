export interface Result<T> {
  status: ResponseStatus;
  data: T;
  error?: Error;
}

export enum ResponseStatus {
  Success = 0,
  InternalError = 1,
  ClientError = 2,
  AuthError = 3
}

export interface Error {
  type: string, // Runtime, KeyNotFound, IndexNotFound
  message: string
}
