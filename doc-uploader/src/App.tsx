import DocumentUploader from "./components/DocumentUploader";

export default function App() {
  return (
    <div className="min-h-screen flex items-center justify-center bg-gray-50 p-4">
      <div className="w-full max-w-md">
      <h1 className="text-3xl font-bold text-center mb-8">Document Uploader</h1>
      <DocumentUploader />
      </div>
    </div>
  );
}
