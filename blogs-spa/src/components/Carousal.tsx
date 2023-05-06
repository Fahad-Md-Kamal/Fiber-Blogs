import { useState } from 'react';
import ImageOne from '../assets/images/nimbus.png';
import ImageTwo from '../assets/images/ai.jpeg';
import ImageThree from '../assets/images/image-1.png';

const images = [ImageOne, ImageTwo, ImageThree];

const Carousel = () => {
  const [currentImage, setCurrentImage] = useState(0);

  const nextImage = () => {
    setCurrentImage((currentImage + 1) % images.length);
  };

  const prevImage = () => {
    setCurrentImage((currentImage + images.length - 1) % images.length);
  };

  return (
    <div className="relative w-full">
      <div className="absolute inset-y-0 left-0 flex items-center">
        <button
          onClick={prevImage}
          className="bg-gray-900 bg-opacity-50 hover:bg-opacity-75 text-white p-2 rounded-l-lg focus:outline-none"
        >
          {/* <FiChevronLeft /> */}
        </button>
      </div>
      <div className="absolute inset-y-0 right-0 flex items-center">
        <button
          onClick={nextImage}
          className="bg-gray-900 bg-opacity-50 hover:bg-opacity-75 text-white p-2 rounded-r-lg focus:outline-none"
        >
          {/* <FiChevronRight /> */}
        </button>
      </div>
      <img
        src={images[currentImage]}
        alt="Carousel image"
        className="object-cover w-full h-64 sm:h-96"
      />
    </div>
  );
};

export default Carousel;
