# Create project with vite + react-ts template
npm create vite@latest doc-uploader -- --template react-ts
cd doc-uploader

# Install dependencies
npm install tailwindcss postcss autoprefixer framer-motion lucide-react clsx tailwind-variants
npm install -D shadcn-ui typescript

# Init Tailwind CSS
npx tailwindcss init -p

# Create folders and files automatically
mkdir -p src/{components/ui,lib,styles,pages}

# Create tailwind config
cat > tailwind.config.ts <<EOF
import type { Config } from "tailwindcss";
const config: Config = {
  content: ["./index.html", "./src/**/*.{ts,tsx}"],
  theme: { extend: {} },
  plugins: [],
};
export default config;
EOF

# Create utils.ts
cat > src/lib/utils.ts <<EOF
import { clsx } from "clsx";
import { twMerge } from "tailwind-variants";

export function cn(...inputs: any[]) {
  return twMerge(clsx(inputs));
}
EOF

# Create globals.css
cat > src/styles/globals.css <<EOF
@tailwind base;
@tailwind components;
@tailwind utilities;
EOF

# Create main.tsx imports globals.css
cat > src/main.tsx <<EOF
import React from "react";
import ReactDOM from "react-dom/client";
import App from "./App";
import "./styles/globals.css";

ReactDOM.createRoot(document.getElementById("root")!).render(
  <React.StrictMode>
    <App />
  </React.StrictMode>
);
EOF

# Create App.tsx with example uploader import
cat > src/App.tsx <<EOF
import DocumentUploader from "@/components/DocumentUploader";

export default function App() {
  return (
    <div className="min-h-screen bg-gray-50 p-8">
      <h1 className="text-3xl font-bold text-center mb-6">Document Uploader</h1>
      <DocumentUploader />
    </div>
  );
}
EOF

# Create DocumentUploader.tsx (paste code from earlier)
# (You can do this manually or automate with echo or a separate script)

echo "Setup done! Now run: npm run dev"
