import React from 'react';
import URLShortener from './components/URLShortener';
import Stats from './components/Stats';
import './App.css';

function App() {
    return (
        <div className="App">
            <header className="App-header">
                <h1>URL Shortener</h1>
            </header>
            <main>
                <URLShortener />
                <Stats />
            </main>
        </div>
    );
}

export default App;
