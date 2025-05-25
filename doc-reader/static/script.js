// Select DOM elements
const uploadArea = document.getElementById('uploadArea');
const browseButton = document.getElementById('browseButton');
const fileInput = document.getElementById('fileInput');
const fileList = document.getElementById('fileList');
const loadingSpinner = document.getElementById('loadingSpinner');

// Handle drag-and-drop functionality
uploadArea.addEventListener('dragover', (event) => {
  event.preventDefault();
  uploadArea.classList.add('drag-over'); // Add a visual cue
});

uploadArea.addEventListener('dragleave', () => {
  uploadArea.classList.remove('drag-over'); // Remove the visual cue
});

uploadArea.addEventListener('drop', (event) => {
  event.preventDefault();
  uploadArea.classList.remove('drag-over');
  const files = event.dataTransfer.files;
  handleFiles(files);
});

// Handle file selection via the "Browse" button
browseButton.addEventListener('click', () => {
  fileInput.click(); // Trigger the hidden file input
});

fileInput.addEventListener('change', () => {
  const files = fileInput.files;
  handleFiles(files);
});

// Function to handle files and display them in the file list
function handleFiles(files) {
  Array.from(files).forEach((file) => {
    const fileItem = document.createElement('div');
    fileItem.classList.add('file-item');

    const fileName = document.createElement('span');
    fileName.textContent = file.name;

    const removeButton = document.createElement('button');
    removeButton.textContent = 'Remove';
    removeButton.addEventListener('click', () => {
      fileItem.remove(); // Remove the file item from the list
    });

    fileItem.appendChild(fileName);
    fileItem.appendChild(removeButton);
    fileList.appendChild(fileItem);
  });
}

browseButton.addEventListener('click', () => {
  fileInput.click();
});

fileInput.addEventListener('change', () => {
  loadingSpinner.hidden = false; // Show spinner
  setTimeout(() => {
    loadingSpinner.hidden = true; // Hide spinner after 2 seconds (simulate upload)
  }, 2000);
});