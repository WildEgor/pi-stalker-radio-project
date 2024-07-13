import L, { IconOptions } from 'leaflet';

interface MarkerIconOptions extends IconOptions {
    iconUrl: string;
}

export class MarkerIcon extends L.Icon<MarkerIconOptions> {
    constructor(options: MarkerIconOptions) {
        super(options);
    }

    createIcon(oldIcon?: HTMLElement): HTMLElement {
        const icon = super.createIcon(oldIcon);
        // Customize the icon element if needed
        return icon;
    }

    createShadow(oldIcon?: HTMLElement): HTMLElement {
        // Customize shadow if needed
        return super.createShadow(oldIcon);
    }
}

// Create an instance of the custom icon with your specific options
export const markerIcon = new MarkerIcon({
    iconUrl: '/path/to/stalker-icon.png', // Path to your custom icon
    iconSize: [32, 32], // Size of the icon
    iconAnchor: [16, 32], // Anchor point of the icon
    popupAnchor: [0, -32], // Popup anchor point
});
