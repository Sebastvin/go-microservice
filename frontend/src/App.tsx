import React, { useState, useRef } from 'react';
import { Upload, Image, Check, Sparkles } from 'lucide-react';

interface StyleOption {
  id: string;
  name: string;
  description: string;
  color: string;
  bgColor: string;
}

const styles: StyleOption[] = [
  { id: 'gta', name: 'GTA', description: 'Grand Theft Auto style', color: 'text-orange-700', bgColor: 'bg-orange-50 border-orange-200' },
  { id: 'retro', name: 'Retro', description: 'Vintage 80s aesthetic', color: 'text-pink-700', bgColor: 'bg-pink-50 border-pink-200' },
  { id: 'anime', name: 'Anime', description: 'Japanese animation style', color: 'text-purple-700', bgColor: 'bg-purple-50 border-purple-200' },
  { id: 'pixel', name: 'Pixel Art', description: '8-bit retro gaming style', color: 'text-green-700', bgColor: 'bg-green-50 border-green-200' },
  { id: 'watercolor', name: 'Watercolor', description: 'Soft artistic painting', color: 'text-blue-700', bgColor: 'bg-blue-50 border-blue-200' },
];

function App() {
  const [selectedImage, setSelectedImage] = useState<File | null>(null);
  const [imagePreview, setImagePreview] = useState<string | null>(null);
  const [selectedStyles, setSelectedStyles] = useState<string[]>([]);
  const [isDragOver, setIsDragOver] = useState(false);
  const fileInputRef = useRef<HTMLInputElement>(null);
  const [isSubmitting, setIsSubmitting] = useState(false);
  const [submitError, setSubmitError] = useState<string | null>(null);

  const handleImageUpload = (file: File) => {
    if (file && file.type.startsWith('image/')) {
      setSelectedImage(file);
      const reader = new FileReader();
      reader.onload = (e) => {
        setImagePreview(e.target?.result as string);
      };
      reader.readAsDataURL(file);
    }
  };

  const handleDragOver = (e: React.DragEvent) => {
    e.preventDefault();
    setIsDragOver(true);
  };

  const handleDragLeave = (e: React.DragEvent) => {
    e.preventDefault();
    setIsDragOver(false);
  };

  const handleDrop = (e: React.DragEvent) => {
    e.preventDefault();
    setIsDragOver(false);
    const files = e.dataTransfer.files;
    if (files.length > 0) {
      handleImageUpload(files[0]);
    }
  };

  const handleFileSelect = (e: React.ChangeEvent<HTMLInputElement>) => {
    const files = e.target.files;
    if (files && files.length > 0) {
      handleImageUpload(files[0]);
    }
  };

  const handleStyleToggle = (styleId: string) => {
    setSelectedStyles(prev => 
      prev.includes(styleId)
        ? prev.filter(id => id !== styleId)
        : [...prev, styleId]
    );
  };

  const isFormValid = selectedImage && selectedStyles.length > 0;

  async function handleSubmit() {
    if (!isFormValid || !selectedImage || !imagePreview) return;
    setIsSubmitting(true);
    setSubmitError(null);
    try {
      const base64 = imagePreview.replace(/^data:image\/(png|jpeg|jpg);base64,/, '');
      const items = selectedStyles.map((style) => ({
        id: "1",
        name: "Onion",
        stylereference: style,
        priceid: "price_1RVZ3hClTXDUG291P1wmsO9h",
      }));
      const uniqueItems = Array.from(
        new Map(items.map(item => [item.id + item.stylereference, item])).values()
      );
      const payload = { items: uniqueItems, image: base64 };
      const res = await fetch('http://localhost:8080/api/customers/3/orders', {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify(payload),
      });

      if (!res.ok) throw new Error('Failed to submit order');
      const data = await res.json();
      if (data.RedirectToURL) {
        window.open(data.RedirectToURL, "_blank");
        return;
      }
      throw new Error('No RedirectToURL in response');
    } catch (err: unknown) {
      if (err instanceof Error) {
        setSubmitError(err.message);
      } else {
        setSubmitError('Unknown error');
      }
    } finally {
      setIsSubmitting(false);
    }
  }

  return (
    <div className="min-h-screen bg-gradient-to-br from-slate-50 to-blue-50 py-8 px-4">
      <div className="max-w-4xl mx-auto">
        {/* Header */}
        <div className="text-center mb-12">
          <div className="flex items-center justify-center mb-4">
            <Sparkles className="w-8 h-8 text-blue-600 mr-3" />
            <h1 className="text-4xl font-bold text-gray-900">Order AI Image Generation</h1>
          </div>
          <p className="text-lg text-gray-600 max-w-2xl mx-auto">
            Transform your images with cutting-edge AI technology. Upload your image and select your preferred artistic style.
          </p>
        </div>

        <div className="bg-white rounded-2xl shadow-xl p-8 space-y-8">
          {/* Image Upload Section */}
          <div className="space-y-6">
            <h2 className="text-2xl font-semibold text-gray-900 flex items-center">
              <Image className="w-6 h-6 mr-2 text-blue-600" />
              Upload Your Image
            </h2>
            
            <div
              className={`relative border-2 border-dashed rounded-xl p-8 text-center transition-all duration-300 cursor-pointer
                ${isDragOver 
                  ? 'border-blue-500 bg-blue-50' 
                  : selectedImage 
                    ? 'border-green-500 bg-green-50' 
                    : 'border-gray-300 bg-gray-50 hover:border-blue-400 hover:bg-blue-50'
                }`}
              onDragOver={handleDragOver}
              onDragLeave={handleDragLeave}
              onDrop={handleDrop}
              onClick={() => fileInputRef.current?.click()}
            >
              <input
                ref={fileInputRef}
                type="file"
                accept="image/*"
                onChange={handleFileSelect}
                className="hidden"
              />
              
              {selectedImage ? (
                <div className="space-y-4">
                  <div className="flex items-center justify-center">
                    <Check className="w-12 h-12 text-green-600" />
                  </div>
                  <div>
                    <p className="text-lg font-medium text-green-700">Image uploaded successfully!</p>
                    <p className="text-sm text-gray-600">{selectedImage.name}</p>
                  </div>
                </div>
              ) : (
                <div className="space-y-4">
                  <div className="flex items-center justify-center">
                    <Upload className="w-12 h-12 text-gray-400" />
                  </div>
                  <div>
                    <p className="text-lg font-medium text-gray-700">
                      {isDragOver ? 'Drop your image here' : 'Drag and drop an image or click to select a file'}
                    </p>
                    <p className="text-sm text-gray-500">Supports JPG, PNG, and other image formats</p>
                  </div>
                </div>
              )}
            </div>

            {/* Image Preview */}
            {imagePreview && (
              <div className="bg-gray-50 rounded-xl p-6">
                <h3 className="text-lg font-medium text-gray-900 mb-4">Image Preview</h3>
                <div className="flex justify-center">
                  <img
                    src={imagePreview}
                    alt="Preview"
                    className="max-w-full max-h-64 rounded-lg shadow-md object-contain"
                  />
                </div>
              </div>
            )}
          </div>

          {/* Style Selection Section */}
          <div className="space-y-6">
            <h2 className="text-2xl font-semibold text-gray-900 flex items-center">
              <Sparkles className="w-6 h-6 mr-2 text-blue-600" />
              Choose Your Style
              {selectedStyles.length > 0 && (
                <span className="ml-2 text-sm bg-blue-100 text-blue-700 px-2 py-1 rounded-full">
                  {selectedStyles.length} selected
                </span>
              )}
            </h2>

            <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-4">
              {styles.map((style) => {
                const isSelected = selectedStyles.includes(style.id);
                return (
                  <div
                    key={style.id}
                    className={`relative border-2 rounded-xl p-6 cursor-pointer transition-all duration-300 hover:shadow-lg
                      ${isSelected 
                        ? `${style.bgColor} border-current ${style.color} shadow-md` 
                        : 'bg-white border-gray-200 hover:border-gray-300 hover:bg-gray-50'
                      }`}
                    onClick={() => handleStyleToggle(style.id)}
                  >
                    {isSelected && (
                      <div className="absolute top-3 right-3">
                        <Check className="w-5 h-5 text-current" />
                      </div>
                    )}
                    <div className="space-y-2">
                      <h3 className="text-lg font-semibold">{style.name}</h3>
                      <p className="text-sm opacity-75">{style.description}</p>
                    </div>
                  </div>
                );
              })}
            </div>
          </div>

          {/* Next Button */}
          <div className="pt-6 border-t border-gray-200">
            <div className="flex flex-col sm:flex-row sm:items-center sm:justify-between gap-4">
              <div className="text-sm text-gray-600">
                {!selectedImage && !selectedStyles.length && (
                  "Please upload an image and select at least one style to continue"
                )}
                {selectedImage && selectedStyles.length === 0 && (
                  "Please select at least one style to continue"
                )}
                {!selectedImage && selectedStyles.length > 0 && (
                  "Please upload an image to continue"
                )}
                {isFormValid && (
                  <span className="text-green-600 font-medium">Ready to proceed!</span>
                )}
              </div>
              <button
                disabled={!isFormValid || isSubmitting}
                className={`px-8 py-3 rounded-xl font-semibold text-white transition-all duration-300
                  ${isFormValid && !isSubmitting
                    ? 'bg-blue-600 hover:bg-blue-700 shadow-lg hover:shadow-xl transform hover:scale-105'
                    : 'bg-gray-300 cursor-not-allowed'
                  }`}
                onClick={handleSubmit}
              >
                {isSubmitting ? 'Processing...' : 'Next Step'}
              </button>
            </div>
            {submitError && (
              <div className="mt-4 text-red-600 font-medium">{submitError}</div>
            )}
          </div>
        </div>
      </div>
    </div>
  );
}

export default App;