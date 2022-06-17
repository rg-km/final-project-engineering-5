import create from 'zustand';
import { devtools, persist } from 'zustand/middleware';

const useAuthStore = create(
  devtools(
    persist(
      (set) => ({
        user: {},
        isAuthenticated: false,
        setUser: (userData) =>
          set(() => ({ user: userData, isAuthenticated: true })),
        logout: () => set(() => ({ user: {}, isAuthenticated: false })),
      }),
      {
        name: 'auth',
      }
    )
  )
);

export default useAuthStore;
