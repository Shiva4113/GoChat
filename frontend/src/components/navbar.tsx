"use client";

import * as React from "react";
import Link from "next/link";
import { Menu } from "lucide-react";
import { ModeToggle } from "./modeToggle";
import { Button } from "./ui/button";

export default function Navbar() {
  const [state, setState] = React.useState(false);

  return (
    <nav className="w-full md:h-[75px] border-b shadow">
      <div className="items-center px-4 max-w-screen-xl mx-auto md:flex md:px-8">
        <div className="flex items-center justify-between py-3 md:py-5 md:block">
          <Link href="/">
            <h1 className="text-3xl font-bold text-primary">GoChat</h1>
          </Link>
          <div className="md:hidden">
            <button
              className="text-gray-700 outline-none p-2 rounded-md focus:border-gray-400 focus:border"
              onClick={() => setState(!state)}
            >
              <Menu />
            </button>
          </div>
        </div>
        <div
          className={`flex-1 justify-self-center pb-3 mt-8 md:block md:pb-0 md:mt-0 ${
            state ? "block" : "hidden"
          }`}
        >
          <ul className="justify-end items-center space-y-8 md:flex md:space-x-6 md:space-y-0">
            <li>
              {/* {status !== "authenticated" && ( */}
              <Link href="/login">
                <Button>Log In</Button>
              </Link>
              {/* )} */}
              {/* {status === "authenticated" && <UserAvatar />} */}
            </li>
            <ModeToggle />
          </ul>
        </div>
      </div>
    </nav>
  );
}
