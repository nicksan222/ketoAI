"use client";

import { cn } from "@/lib/utils";
import { Button } from "../ui/button";

import { FiClock, FiHeart, FiList, FiPlusCircle, FiShare, FiUser } from "react-icons/fi";
import SidebarSection from "./SidebarSection";
import SidebarButton from "./SidebarButton";

interface Props {
  children: React.ReactNode;
  className?: string;
}

export function Sidebar({ className, children }: Props) {
  return (
    <>
      <button
        data-drawer-target="default-sidebar"
        data-drawer-toggle="default-sidebar"
        aria-controls="default-sidebar"
        type="button"
        className="inline-flex items-center p-2 mt-2 ml-3 text-sm text-gray-500 rounded-lg sm:hidden hover:bg-gray-100 focus:outline-none focus:ring-2 focus:ring-gray-200 dark:text-gray-400 dark:hover:bg-gray-700 dark:focus:ring-gray-600"
      >
        <span className="sr-only">Apri sidebar</span>
        <svg
          className="w-6 h-6"
          aria-hidden="true"
          fill="currentColor"
          viewBox="0 0 20 20"
          xmlns="http://www.w3.org/2000/svg"
        >
          <path
            clipRule="evenodd"
            fillRule="evenodd"
            d="M2 4.75A.75.75 0 012.75 4h14.5a.75.75 0 010 1.5H2.75A.75.75 0 012 4.75zm0 10.5a.75.75 0 01.75-.75h7.5a.75.75 0 010 1.5h-7.5a.75.75 0 01-.75-.75zM2 10a.75.75 0 01.75-.75h14.5a.75.75 0 010 1.5H2.75A.75.75 0 012 10z"
          ></path>
        </svg>
      </button>

      <aside
        id="default-sidebar"
        className="fixed top-0 left-0 z-40 w-64 h-screen transition-transform -translate-x-full sm:translate-x-0"
        aria-label="Sidebar"
      >
        <div className="h-full px-3 py-4 overflow-y-auto bg-gray-50 dark:bg-gray-800">
          <div className={cn("pb-12", className)}>
            <SidebarSection text="Ricette">
              <SidebarButton
                icon={<FiHeart />}
                text="Preferiti"
                goTo="/recipes/favorites"
              />
              <SidebarButton
                icon={<FiList />}
                text="Scopri"
                goTo="/recipes/discover"
              />
            </SidebarSection>
            <SidebarSection text="Ingredienti">
              <SidebarButton
                icon={<FiHeart />}
                text="Preferenze"
                goTo="/ingredients/favorites"
              />
            </SidebarSection>
            <SidebarSection text="Crea">
              <SidebarButton icon={<FiPlusCircle />} text="Ricetta" goTo="/recipes/new"/>
              <SidebarButton icon={<FiClock />} text="In approvazione" goTo="/recipes/waiting" />
            </SidebarSection>
            <SidebarSection text="Impostazioni">
              <SidebarButton icon={<FiUser />} text="Account" />
            </SidebarSection>
          </div>
        </div>
      </aside>

      <div className="md:p-8 p-4 sm:ml-64">{children}</div>
    </>
  );
}
