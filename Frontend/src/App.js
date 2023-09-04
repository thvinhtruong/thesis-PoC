
import './App.css';
import {
    BrowserRouter as Router,
    Routes,
    Route,
    Link
} from "react-router-dom";
import About from "./pages/About";
import Home from "./pages/Home";
import Register from "./pages/Register";


function App() {
    return (
        <Router>
            <div>
                <nav>
                    <ul>
                        <li>
                            <Link to="/">Home</Link>
                        </li>
                        <li>
                            <Link to="/about">About</Link>
                        </li>
                        <li>
                            <Link to="/register">Register</Link>
                        </li>
                    </ul>
                </nav>
                <Routes>
                    <Route path="/about" element={<About/>}/>
                    <Route path="/register" element={<Register/>}/>
                    <Route path="/" element={<Home/>}/>
                </Routes>
            </div>
        </Router>
    );
}

export default App;
