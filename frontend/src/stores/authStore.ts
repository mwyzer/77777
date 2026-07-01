import { create } from "zustand";
import type { User } from "../types";
import { login as apiLogin, getMe } from "../services/api";

interface AuthState {
  token: string | null;
  user: User | null;
  isLoading: boolean;
  error: string | null;

  login: (email: string, password: string) => Promise<void>;
  fetchMe: () => Promise<void>;
  logout: () => void;
  clearError: () => void;
}

export const useAuthStore = create<AuthState>((set) => ({
  token: localStorage.getItem("token"),
  user: null,
  isLoading: false,
  error: null,

  login: async (email, password) => {
    set({ isLoading: true, error: null });
    try {
      const data = await apiLogin({ email, password });
      localStorage.setItem("token", data.token);
      set({ token: data.token, user: data.user, isLoading: false });
    } catch {
      set({ isLoading: false, error: "Invalid email or password" });
      throw new Error("Login failed");
    }
  },

  fetchMe: async () => {
    try {
      const user = await getMe();
      set({ user });
    } catch {
      set({ token: null, user: null });
      localStorage.removeItem("token");
    }
  },

  logout: () => {
    localStorage.removeItem("token");
    set({ token: null, user: null });
  },

  clearError: () => set({ error: null }),
}));
