import { Outlet } from 'react-router-dom';
import Navigation from './components/Navigation';

function App() {
  return (
    <>
      <Navigation />
      <main>
        <Outlet />
      </main>
    </>
  );
}

export default App;
