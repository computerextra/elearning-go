import { House, Info, School } from "lucide-react";
import { NavLink, Outlet } from "react-router";
import ThemeSelector from "./Components/ThemeSelector";

export default function Layout() {
  return (
    <div className="bg-base-100 max-w-[100rem] mx-auto">
      <div className="navbar bg-base-100 shadow-sm mb-5">
        <div className="navbar-start">
          <NavLink to="/" className="btn btn-ghost text-xl">
            Datenschutz Training
          </NavLink>
        </div>
        <div className="navbar-end">
          <ThemeSelector />
          {/* Profile Badge */}
          <div className="dropdown dropdown-end">
            <div
              tabIndex={0}
              role="button"
              className="btn btn-ghost btn-circle avatar"
            >
              <div className="w-10 rounded-full">
                <img
                  alt="Tailwind CSS Navbar component"
                  src="https://img.daisyui.com/images/stock/photo-1534528741775-53994a69daeb.webp"
                />
              </div>
            </div>
            <ul
              tabIndex={0}
              className="menu menu-sm dropdown-content bg-base-100 rounded-box z-1 mt-3 w-52 p-2 shadow"
            >
              <li>
                <a className="justify-between">
                  Profile
                  <span className="badge">New</span>
                </a>
              </li>
              <li>
                <a>Settings</a>
              </li>
              <li>
                <a>Logout</a>
              </li>
            </ul>
          </div>
        </div>
      </div>

      <Outlet />

      <div className="dock">
        <NavLink
          to="/"
          className={({ isActive }) => (isActive ? "dock-active" : "")}
        >
          <House />
          <span className="dock-label">Startseite</span>
        </NavLink>

        <NavLink to="/Info">
          <Info />
          <span className="dock-label">Infos</span>
        </NavLink>

        <NavLink to="/Settings">
          <School />
          <span className="dock-label">Kurse</span>
        </NavLink>

        <NavLink to="/Impressum">
          <School />
          <span className="dock-label">Impressum</span>
        </NavLink>

        <NavLink to="/Datenschutz">
          <School />
          <span className="dock-label">Datenschutz</span>
        </NavLink>
      </div>
    </div>
  );
}
