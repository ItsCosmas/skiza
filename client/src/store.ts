import { create } from 'zustand'


// Define the state interface
interface StoreState {
  messages: any[]; // Use a more specific type if you know the shape of the message
  addToMessages: (message: any) => void; // Change 'any' to a more specific type if possible
  isConnected: boolean;
  updateIsConnected: (shouldConnect: boolean) => void;
  shouldConnect: boolean;
  updateShouldConnect: (shouldConnect: boolean) => void;
}

export const useStore = create<StoreState>((set) => ({
  messages: [],
  addToMessages: (message: any) =>
    set((state: { messages: any }) => ({
      messages: [...state.messages, message],
    })),
    shouldConnect: false,
    updateShouldConnect: (shouldConnect: boolean) => 
      set(() => ({ shouldConnect })),
    isConnected: false,
    updateIsConnected: (isConnected: boolean) => 
      set(() => ({ isConnected })),
}));



