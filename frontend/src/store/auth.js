import create from 'zustand';

const useAuthStore = create((set) => ({
  user: {},
  isAuthenticated: false,
  setUser: (userData) => set(() => ({ user: userData, isAuthenticated: true })),
}));

export default useAuthStore;
