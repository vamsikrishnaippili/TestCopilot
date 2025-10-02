import React, { useState } from "react";

const VideoDownloader = () => {
  const [link, setLink] = useState("");
  const [formats, setFormats] = useState([]);
  const [selectedFormat, setSelectedFormat] = useState("");
  const [downloadUrl, setDownloadUrl] = useState("");
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState("");

  const handleFetchFormats = async () => {
    setLoading(true);
    setError("");
    setDownloadUrl("");
    try {
      const res = await fetch("https://testcopilot.onrender.com/api/fetch-formats", {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify({ link }),
      });
      const data = await res.json();
      if (data.success) {
        setFormats(data.formats);
        setSelectedFormat(data.formats[0]);
      } else {
        setError(data.message || "Failed to fetch formats");
      }
    } catch (err) {
      setError("Failed to process the link.");
    }
    setLoading(false);
  };

  const handleDownload = () => {
    const url = `https://testcopilot.onrender.com/api/download?link=${encodeURIComponent(link)}&format=${encodeURIComponent(selectedFormat)}`;
    window.open(url, "_blank");
  };

  return (
    <div className="max-w-lg mx-auto p-4 border rounded shadow bg-white">
      <h1 className="text-2xl font-bold mb-2">TeraBox Video Downloader</h1>
      <input
        className="w-full p-2 border rounded mb-2"
        type="text"
        placeholder="Paste TeraBox video link here"
        value={link}
        onChange={e => setLink(e.target.value)}
      />
      <button
        className="bg-blue-500 text-white px-4 py-2 rounded mb-2 w-full"
        onClick={handleFetchFormats}
        disabled={loading || !link}
      >
        {loading ? "Processing..." : "Fetch Video"}
      </button>

      {formats.length > 0 && (
        <div className="mb-2">
          <label>Select Format: </label>
          <select
            className="border rounded p-1"
            value={selectedFormat}
            onChange={e => setSelectedFormat(e.target.value)}
          >
            {formats.map(f => (
              <option key={f} value={f}>{f}</option>
            ))}
          </select>
        </div>
      )}

      {formats.length > 0 && (
        <button
          className="bg-green-500 text-white px-4 py-2 rounded w-full"
          onClick={handleDownload}
          disabled={loading}
        >
          Download Now
        </button>
      )}

      {error && (
        <div className="text-red-600 mt-2">{error}</div>
      )}
    </div>
  );
};

export default VideoDownloader;  
