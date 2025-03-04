import {createContext, ReactNode, useContext, useState} from "react";

type NavContextType = {
  open: boolean;
  setOpen: (open: boolean) => void;
};

const NavContext = createContext<NavContextType>({
  open: false,
  setOpen: () => null,
});

export function NavProvider({ children }: { children: ReactNode }) {
  const [open, setOpen] = useState(false);

  return (
    <NavContext.Provider value={{ open, setOpen }}>
      {children}
    </NavContext.Provider>
  );
}

export function useNav() {
  const context = useContext(NavContext);

  if (context === undefined) {
    throw new Error("useNav must be used within a NavProvider");
  }

  return context;
}
