import React, { useState } from 'react';
import '../App.css';

const URLShortener = () => {
    const [longURL, setLongURL] = useState('');
    const [shortURL, setShortURL] = useState('');

    const handleSubmit = async (e) => {
        e.preventDefault();
        const response = await fetch('http://localhost:8080/shorten', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify({ long_url: longURL })
        });
        const data = await response.json();
        setShortURL(data.short_url);
    };

    return (
        <div className="url-shortener">
            <h2>URL Shortener</h2>
            <form onSubmit={handleSubmit}>
                <input
                    type="text"
                    value={longURL}
                    onChange={(e) => setLongURL(e.target.value)}
                    placeholder="Enter long URL"
                    required
                />
                <button type="submit">Shorten</button>
            </form>
            {shortURL && (
                <div className="result">
                    <p>Short URL: <a href={shortURL} target="_blank" rel="noopener noreferrer">{shortURL}</a></p>
                </div>
            )}
        </div>
    );
};

export default URLShortener;
