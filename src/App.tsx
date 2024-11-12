import { lazy } from 'react';
import { BrowserRouter, Route, Routes } from 'react-router-dom';
import { PATH } from './constants/paths';

const Home = lazy(() => import('./pages/Home'))

function App() {
  return (
    <BrowserRouter basename={process.env.PUBLIC_URL}>
      <Routes>
        <Route path={PATH.HOME} Component={Home}></Route>
      </Routes>
    </BrowserRouter>
  );
}

export default App;
