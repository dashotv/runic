import axios from "axios";
import { useEffect, useState } from "react";
import "./App.css";

function App() {
  const [releases, setReleases] = useState<Release[]>([]);

  useEffect(() => {
    axios.get("/api/releases/?limit=100").then((response) => {
      setReleases(response.data?.results);
    });
  }, []);

  return (
    <>
      <h3>Runic</h3>
      {releases.map((release) => (
        <div key={release.id}>
          {release.source} {release.title || release.raw.title}{" "}
          {release.year > 0 ? `(${release.year})` : ""} {release.season}x
          {release.episode} [{release.group}/{release.website}] {release.size}B
        </div>
      ))}
    </>
  );
}

export interface Release {
  id: string;
  title: string;
  type: string;
  source: string;
  year: number;
  description: string;
  size: number;
  download: string;
  infohash: string;
  season: number;
  episode: number;
  volume: number;
  group: string;
  website: string;
  verified: boolean;
  widescreen: boolean;
  unrated: boolean;
  uncensored: boolean;
  bluray: boolean;
  threeD: boolean;
  resolution: string;
  encodings: string[];
  quality: string;
  raw: {
    title: string;
  };

  created_at: string;
  updated_at: string;
}

export default App;
