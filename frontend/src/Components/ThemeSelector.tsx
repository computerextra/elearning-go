import { useEffect } from "react";
import { themeChange } from "theme-change";

const themes = [
  {
    name: "",
    title: "System",
  },
  {
    name: "light",
    title: "Hell",
  },
  {
    name: "dark",
    title: "Dunkel",
  },
  {
    name: "cupcake",
    title: "Cupcake",
  },
  {
    name: "bumblebee",
    title: "Bumblebee",
  },
  {
    name: "emerald",
    title: "Emerald",
  },
  {
    name: "corporate",
    title: "corporate",
  },
  {
    name: "synthwave",
    title: "synthwave",
  },
  {
    name: "retro",
    title: "retro",
  },
  {
    name: "cyberpunk",
    title: "cyberpunk",
  },
  {
    name: "valentine",
    title: "valentine",
  },
  {
    name: "halloween",
    title: "halloween",
  },
  {
    name: "garden",
    title: "garden",
  },
  {
    name: "forest",
    title: "forest",
  },
  {
    name: "aqua",
    title: "aqua",
  },
  {
    name: "lofi",
    title: "lofi",
  },
  {
    name: "pastel",
    title: "pastel",
  },
  {
    name: "fantasy",
    title: "fantasy",
  },
  {
    name: "wireframe",
    title: "wireframe",
  },
  {
    name: "black",
    title: "black",
  },
  {
    name: "luxury",
    title: "luxury",
  },
  {
    name: "dracula",
    title: "dracula",
  },
  {
    name: "cmyk",
    title: "cmyk",
  },
  {
    name: "autumn",
    title: "autumn",
  },
  {
    name: "business",
    title: "business",
  },
  {
    name: "acid",
    title: "acid",
  },
  {
    name: "lemonade",
    title: "lemonade",
  },
  {
    name: "night",
    title: "night",
  },
  {
    name: "coffee",
    title: "coffee",
  },
  {
    name: "winter",
    title: "winter",
  },
  {
    name: "dim",
    title: "dim",
  },
  {
    name: "nord",
    title: "nord",
  },
  {
    name: "sunset",
    title: "sunset",
  },
  {
    name: "caramellatte",
    title: "caramellatte",
  },
  {
    name: "abyss",
    title: "abyss",
  },
  {
    name: "silk",
    title: "silk",
  },
];

export default function ThemeSelector() {
  useEffect(() => {
    themeChange(false);
    // 👆 false parameter is required for react project
  }, []);

  return (
    <div className="dropdown dropdown-end block">
      <div tabIndex={0} role="button" className="btn btn-sm btn-ghost">
        <svg
          width="20"
          height="20"
          xmlns="http://www.w3.org/2000/svg"
          fill="none"
          viewBox="0 0 24 24"
          className="h-5 w-5 stroke-current md:hidden"
        >
          <path
            stroke-linecap="round"
            stroke-linejoin="round"
            stroke-width="2"
            d="M7 21a4 4 0 01-4-4V5a2 2 0 012-2h4a2 2 0 012 2v12a4 4 0 01-4 4zm0 0h12a2 2 0 002-2v-4a2 2 0 00-2-2h-2.343M11 7.343l1.657-1.657a2 2 0 012.828 0l2.829 2.829a2 2 0 010 2.828l-8.486 8.485M7 17h.01"
          ></path>
        </svg>
        <span className="hidden font-normal md:inline">Theme</span>
        <svg
          width="12px"
          height="12px"
          className="hidden h-2 w-2 fill-current opacity-60 sm:inline-block"
          xmlns="http://www.w3.org/2000/svg"
          viewBox="0 0 2048 2048"
        >
          <path d="M1799 349l242 241-1017 1017L7 590l242-241 775 775 775-775z"></path>
        </svg>
      </div>
      <div
        tabIndex={0}
        className="dropdown-content bg-base-200 text-base-content rounded-box top-px h-[28rem] max-h-[calc(100vh-10rem)] overflow-y-auto border border-white/5 shadow-2xl outline-1 outline-black/5 mt-16"
      >
        <ul className="menu w-56">
          {themes.map((theme) => (
            <li key={theme.title}>
              <button
                className="px-2 gap-3"
                data-set-theme={theme.name}
                data-act-class="outline-base-content!"
              >
                <div
                  data-theme={theme.name}
                  className="grid grid-cols-2 gap-0.5 p-1 rounded-md shadow-sm shrink-0 bg-base-100"
                >
                  <div className="size-1 rounded-full bg-base-content"></div>{" "}
                  <div className="size-1 rounded-full bg-primary"></div>{" "}
                  <div className="size-1 rounded-full bg-secondary"></div>{" "}
                  <div className="size-1 rounded-full bg-accent"></div>
                </div>{" "}
                <div className="w-32 truncate">{theme.title.toUpperCase()}</div>{" "}
                <svg
                  xmlns="http://www.w3.org/2000/svg"
                  width="16"
                  height="16"
                  viewBox="0 0 24 24"
                  fill="currentColor"
                  className="invisible h-3 w-3 shrink-0"
                >
                  <path d="M20.285 2l-11.285 11.567-5.286-5.011-3.714 3.716 9 8.728 15-15.285z"></path>
                </svg>
              </button>
            </li>
          ))}
        </ul>
      </div>
    </div>
  );
}
