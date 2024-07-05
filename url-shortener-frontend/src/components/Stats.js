import React, { useEffect, useState } from 'react';
import '../App.css';

const Stats = () => {
    const [stats, setStats] = useState([]);

    useEffect(() => {
        const fetchStats = async () => {
            const response = await fetch('http://localhost:8080/stats');
            const data = await response.json();
            setStats(data);
        };
        fetchStats();
    }, []);

    return (
        <div className="stats">
            <h2>URL Statistics</h2>
            <ul>
                {stats.map((url) => (
                    <li key={url.id}>
                        <p>Long URL: {url.long_url}</p>
                        <p>Short URL: <a href={url.short_url} target="_blank" rel="noopener noreferrer">{url.short_url}</a></p>
                        <p>Clicks: {url.clicks}</p>
                    </li>
                ))}
            </ul>
        </div>
    );
};

export default Stats;
