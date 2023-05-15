import { BrowserRouter, Routes, Route } from 'react-router-dom';
import Navbar from './components/Navbar';
import Login from './components/Login';
import { ProblemSet } from './components/ProblemSet';
import SingleProblem from './components/SingleProblem';
import SignUp from './components/SignUp';
import Home from './components/Home';

function App() {

    /* Add routing here, routes look like -
        /login - Login page
        /signup - Signup page
        /problemset/all/ - All problems (see problems array above)
        /problems/:problem_slug - A single problem page
     */


    return (
        <>
            <BrowserRouter>
                <Navbar />
                <Routes>
                    <Route path="/" element={<Home />} />
                    <Route path="/login" element={<Login />} />
                    <Route path="/signup" element={<SignUp />} />
                    <Route path="/problems/all" element={<ProblemSet />} />
                    <Route path="/problems/:pid" element={<SingleProblem />} />

                </Routes>
            </BrowserRouter>
        </>
    )
}

export default App;
