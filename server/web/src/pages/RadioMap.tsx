import "leaflet/dist/leaflet.css";
import { MapContainer, Marker, Popup, TileLayer, CircleMarker } from "react-leaflet";
import MarkerClusterGroup from "react-leaflet-cluster";
import { Loader2, Menu, Pause, Play } from "lucide-react";
import axios from "axios";
import { useEffect, useState } from "react";
import { StationLocation } from "@/types/types.ts";
import { Button } from "@/atoms/button.tsx";
import { useAudio } from "@/store/providers/audioContext.tsx";
import { useNav } from "@/store/providers/navContext.tsx";
import {DivIcon, LatLngTuple} from "leaflet";
import { divIcon, point } from "leaflet";

const createClusterCustomIcon = function (cluster: any): DivIcon {
    return divIcon({
        html: `<span class="leaflet-cluster-icon">${cluster.getChildCount()}</span>`,
        className: "custom-marker-cluster",
        iconSize: point(33, 33, true)
    });
};

const RadioMap = () => {
    const { track, setTrack, sound, playing, setLoadingAudio } = useAudio();
    const { setOpen } = useNav();

    const [data, setData] = useState<StationLocation[]>([]);
    const [loading, setLoading] = useState(false);

    const [center, setCenter] = useState<LatLngTuple>([84, -75.524426]); // Default center

    const handlePlay = () => {
        setLoadingAudio(true);
        sound?.play();
    };

    const handleSetTrack = (prop: StationLocation) => {
        setLoadingAudio(true);
        setTrack(prop as any);
    };

    useEffect(() => {
        setLoading(true);
        axios
            .get("https://de1.api.radio-browser.info/json/stations/search?has_geo_info=true&limit=500&hidebroken=true")
            .then((res) => {
                setData(res.data);
            })
            .catch((err) => console.log(err))
            .finally(() => setLoading(false));
    }, []);

    // Define the bounds
    const maxBounds: LatLngTuple[] = [
        [84, -125], // Southwest corner
        [-72, -125]  // Northeast corner
    ];

    // from left top corner in clockwise dir
    const cornerBounds: LatLngTuple[] = [
        [84, -170.524426],
        [84, -75.524426],
        [-72, -75.524426],
        [-72, -170.524426],
    ]

    return (
        <div className="h-screen w-screen overflow-hidden relative">
            <div className="py-6 text-center">
                <h1 className="text-5xl font-bold">Stalker</h1>
            </div>
            <Button
                variant={"unstyled"}
                onClick={() => setOpen(true)}
                className="absolute left-4 top-4 h-16 w-16 lg:hidden text-foreground"
            >
                <Menu className="w-full h-full" />
            </Button>
            {loading ? (
                <Loader2 className="mx-auto my-10 animate-spin w-16 h-16" />
            ) : (
                <MapContainer
                    whenReady={() => setCenter(center)}
                    className="h-full w-full"
                    center={center}
                    zoom={4}
                    minZoom={4}
                    maxZoom={7}
                    trackResize={true}
                    fadeAnimation={true}
                    zoomAnimation={true}
                    keyboard={true}
                    scrollWheelZoom={true}
                    dragging={true}
                    maxBounds={[
                        maxBounds[0] as LatLngTuple, // Southwest corner
                        maxBounds[1] as LatLngTuple  // Northeast corner
                    ]}
                    maxBoundsViscosity={1.0}
                >
                    <TileLayer url={`/map/{z}/{x}/{y}.png`} />
                    <MarkerClusterGroup
                        chunkedLoading
                        iconCreateFunction={createClusterCustomIcon}
                    >
                        {data?.map((item, index) => (
                            <Marker
                                // icon={markerIcon}
                                key={index}
                                position={[item.geo_lat, item.geo_long]}>
                                <Popup>
                                    <div className="flex flex-col justify-center items-center w-40 gap-1 text-center">
                                        <img src={item.favicon} className="w-full" alt="" />
                                        <h1 className="font-semibold">{item.name}</h1>
                                        <h1>{item.country}</h1>
                                        {track?.url === item.url ? (
                                            playing ? (
                                                <Button
                                                    onClick={() => sound?.pause()}
                                                    className="shadow-lg p-3 h-auto w-auto rounded-full"
                                                >
                                                    <Pause className="" />
                                                </Button>
                                            ) : (
                                                <Button
                                                    onClick={handlePlay}
                                                    className="shadow-lg p-3 h-auto w-auto rounded-full"
                                                >
                                                    <Play className="" />
                                                </Button>
                                            )
                                        ) : (
                                            <Button
                                                onClick={() => handleSetTrack(item)}
                                                className="shadow-lg p-3 h-auto w-auto rounded-full"
                                            >
                                                <Play className="" />
                                            </Button>
                                        )}
                                    </div>
                                </Popup>
                            </Marker>
                        ))}
                    </MarkerClusterGroup>

                    {/* Add corners */}
                    <CircleMarker center={cornerBounds[0]} radius={10} color="red" />
                    <CircleMarker center={cornerBounds[1]} radius={10} color="red" />
                    <CircleMarker center={cornerBounds[2]} radius={10} color="red" />
                    <CircleMarker center={cornerBounds[3]} radius={10} color="red" />

                    {/* Add CircleMarkers to show maxBounds */}
                    <CircleMarker center={maxBounds[0]} radius={5} color="red" />
                    <CircleMarker center={maxBounds[1]} radius={5} color="blue" />
                </MapContainer>
            )}
        </div>
    );
};

export default RadioMap;
