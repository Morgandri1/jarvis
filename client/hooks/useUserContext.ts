import { useContext, createContext, ReactNode, useState } from "react";
import User from "@/constants/User";

interface UserContextInner {
  user: User;
  setUser: (user: User) => void;
}

export const UserContext = createContext<UserContextInner>(null);

export default function useUserContext() {
  const context = useContext(UserContext);
  if (!context) {
    throw new Error("useUserContext must be used within a UserProvider");
  }
  return context;
}

export function UserProvider({ children, _user }: { children: ReactNode, _user?: any }) {
  const [user, setUser] = useState<User>(_user);

  return (
    <UserContext.Provider value={{ user, setUser }}>
      {children}
    </UserContext.Provider>
  );
}