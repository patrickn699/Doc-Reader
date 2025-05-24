import DocumentUploader from "@/components/DocumentUploader";

export default function App() {
  return (
    <div className="min-h-screen bg-gray-50 p-8">
      <h1 className="text-3xl font-bold text-center mb-6">Document Uploader</h1>
      <DocumentUploader />
    </div>
  );
}
