import React, { useState } from "react";
import { Card, CardContent } from "../components/ui/card";
import { Button } from "../components/ui/button";
import { UploadCloud, FileText, Loader2 } from "lucide-react";
import { motion } from "framer-motion";
import { cn } from "../lib/utils";

export default function DocumentUploader() {
  const [files, setFiles] = useState<File[]>([]);
  const [loading, setLoading] = useState(false);

  const handleFileChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    const selectedFiles = e.target.files ? Array.from(e.target.files) : [];
    setFiles(selectedFiles);
  };

  const handleUpload = async () => {
    if (files.length === 0) return;

    setLoading(true);

    // Simulate upload delay
    setTimeout(() => {
      alert(`Uploaded ${files.length} file(s)!`);
      setFiles([]);
      setLoading(false);
    }, 2000);
  };

  return (
    <motion.div
      className="max-w-xl mx-auto p-6 space-y-6"
      initial={{ opacity: 0, y: 30 }}
      animate={{ opacity: 1, y: 0 }}
      transition={{ duration: 0.5 }}
    >
      <Card className="border-dashed border-2 border-gray-300 hover:border-blue-500 transition cursor-pointer">
        <CardContent className="flex flex-col items-center space-y-4">
          <UploadCloud className="h-12 w-12 text-blue-500" />
          <input
            type="file"
            id="file-upload"
            multiple
            onChange={handleFileChange}
            className="hidden"
          />
          <label
            htmlFor="file-upload"
            className="text-blue-600 hover:underline cursor-pointer select-none"
          >
            Click here to select documents
          </label>

          {files.length > 0 && (
            <div className="w-full space-y-2 max-h-48 overflow-auto">
              {files.map((file, idx) => (
                <div
                  key={idx}
                  className="flex items-center justify-between bg-gray-100 rounded px-3 py-2"
                >
                  <div className="flex items-center gap-2 text-gray-700">
                    <FileText className="h-5 w-5" />
                    <span className="truncate max-w-xs">{file.name}</span>
                  </div>
                  <span className="text-xs text-gray-500">
                    {(file.size / 1024).toFixed(1)} KB
                  </span>
                </div>
              ))}
            </div>
          )}
        </CardContent>
      </Card>

      <Button
        onClick={handleUpload}
        disabled={loading || files.length === 0}
        className={cn("w-full", loading && "opacity-60 cursor-not-allowed")}
      >
        {loading && <Loader2 className="mr-2 h-4 w-4 animate-spin" />}
        {loading ? "Uploading..." : "Upload"}
      </Button>
    </motion.div>
  );
}
