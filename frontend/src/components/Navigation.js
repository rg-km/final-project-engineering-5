import { Link } from 'react-router-dom';

function Navigation() {
  return (
    <header>
      <div className="container mx-auto flex items-center px-10 py-6">
        <h1 className="text-xl font-semibold">Sistem Informasi Beasiswa</h1>
        <nav className="ml-auto space-x-8">
          <Link to="/" className="hover:underline">
            Beranda
          </Link>
          <Link to="/beasiswa" className="hover:underline">
            Beasiswa
          </Link>
        </nav>
        <Link
          to="/login"
          className="ml-10 rounded border border-black px-4 py-1 hover:bg-gray-200"
        >
          Login
        </Link>
        <Link
          to="/register"
          className="ml-4 rounded border border-transparent bg-black px-4 py-1 text-white hover:bg-gray-800"
        >
          Register
        </Link>
      </div>
    </header>
  );
}

export default Navigation;
