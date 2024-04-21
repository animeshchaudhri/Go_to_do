import Link from "next/link";
import React from "react";
import { BsGithub } from "react-icons/bs";

const Footer = () => {
  return (
    <footer className="flex items-center justify-center text-gray-400 fixed bottom-0 w-full h-16 z-0">
      <div className="max-w-3xl mx-auto py-4 px-5 flex flex-col items-center justify-between md:flex-row">
        <p className="text-sm  text-center md:text-left">
          Developed by animesh chaudhri with Next.js. All rights reserved.
          &copy; {new Date().getFullYear()}
        </p>
        <Link href="https://github.com/animeshchaudhri">
          <div className="social-media mt-4 md:mt-0">
            <BsGithub size={24} className="ml-2" />
          </div>
        </Link>
      </div>
    </footer>
  );
};

export default Footer;
