import { FC, ReactNode, createContext, useContext, useState } from "react";
import { User } from "../types/authentication";

// interface User {
//   username: string;
//   email: string;
// }

interface AuthContext {
  isLogged: boolean;
  user?: User;

  updateUser: (user?: User) => void;
}

const AuthContext = createContext<AuthContext>({
  isLogged: false,
  updateUser: () => {},
});
export const useAuth = () => useContext(AuthContext);

type AuthProviderProps = {
  children: ReactNode;
};

const AuthProvider: FC<AuthProviderProps> = ({ children }) => {
  const [isLoggedIn, setIsLoggedIn] = useState(false);
  const [user, setUser] = useState<User>();

  const updateUser = (user?: User) => {
    if (user) {
      setIsLoggedIn(true);
    }
    setUser(user);
  };

  return (
    <AuthContext.Provider value={{ isLogged: isLoggedIn, user, updateUser }}>
      {children}
    </AuthContext.Provider>
  );
};

export default AuthProvider;
