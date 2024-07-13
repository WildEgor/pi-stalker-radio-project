import {createContext, Dispatch, ReactNode, SetStateAction, useContext, useEffect, useState} from "react";
import { Howl } from "howler";
import { Station } from "@/types/types";
import { useToast } from "@/hooks/use-toast";

type AudioContextType = {
  track: Station | undefined;
  setTrack: Dispatch<SetStateAction<Station | undefined>>;
  sound: Howl | null | undefined;
  playing: boolean;
  setPlaying: Dispatch<SetStateAction<boolean>>;
  volume: number;
  setVolume: Dispatch<SetStateAction<number>>;
  muted: boolean;
  setMuted: Dispatch<SetStateAction<boolean>>;
  likedStations: Station[];
  setLikedStations: Dispatch<SetStateAction<Station[]>>;
  handleLike: (station: Station) => void;
  loadingAudio: boolean;
  setLoadingAudio: Dispatch<SetStateAction<boolean>>;
};

const AudioContext = createContext<AudioContextType>({
  track: {
    country: "",
    favicon: "",
    language: "",
    name: "",
    tags: "",
    url: "",
    votes: 0,
    serveruuid: "",
    stationuuid: "",
  },
  setTrack: () => null,
  sound: null,
  playing: false,
  setPlaying: () => null,
  volume: 0.5,
  setVolume: () => null,
  muted: false,
  setMuted: () => null,
  likedStations: [],
  setLikedStations: () => null,
  handleLike: () => null,
  loadingAudio: false,
  setLoadingAudio: () => null,
});

export function AudioProvider({ children }: { children: ReactNode }) {
  const { toast } = useToast();

  const [track, setTrack] = useState<Station>();
  const [sound, setSound] = useState<Howl | null>(null);
  const [playing, setPlaying] = useState(false);
  const [volume, setVolume] = useState(0.5);
  const [muted, setMuted] = useState(false);
  const [likedStations, setLikedStations] = useState<Station[]>([]);
  const [loadingAudio, setLoadingAudio] = useState(false);

  const handleLike = (station: Station) => {
    const liked = localStorage.getItem("liked");

    if (liked) {
      const likedStations = JSON.parse(liked);
      const stationExists = likedStations.find(
        (likedStation: Station) =>
          likedStation.stationuuid === station.stationuuid
      );

      if (stationExists) {
        const newLikedStations = likedStations.filter(
          (likedStation: Station) => likedStation.url !== station.url
        );
        localStorage.setItem("liked", JSON.stringify(newLikedStations));
        setLikedStations(newLikedStations);
        toast({
          description: "Station removed from favorites",
        });
      } else {
        const newLikedStations = [...likedStations, station];
        localStorage.setItem("liked", JSON.stringify(newLikedStations));
        setLikedStations(newLikedStations);
        toast({
          description: "Station added to favorites",
        });
      }
    } else {
      localStorage.setItem("liked", JSON.stringify([station]));
      setLikedStations([station]);
      toast({
        description: "Station added to favorites",
      });
    }
  };

  useEffect(() => {
    const liked = localStorage.getItem("liked");
    setLikedStations(liked ? JSON.parse(liked) : []);

    if (!track) return;

    const newSound = new Howl({
      src: [track.url],
      html5: true,
      onload: () => {
        console.log("loaded");

        setLoadingAudio(false);
      },
      onplay: () => {
        console.log("playing");

        setPlaying(true);
        setLoadingAudio(false);
      },
      onpause: () => {
        console.log("paused");

        setPlaying(false);
        setLoadingAudio(false);
      },
      onstop: () => {
        setPlaying(false);
      },
      onvolume: () => {
        // console.log(newSound?.volume());
        // if (newSound?.volume() === 0) {
        //   setMuted(true);
        // } else {
        //   setMuted(false);
        // }
      },
      onmute: () => {},
    });

    sound?.stop();
    setSound(newSound);
    newSound?.play();

    newSound.volume(volume);
  }, [track]);

  return (
    <AudioContext.Provider
      value={{
        likedStations,
        setLikedStations,
        track,
        setTrack,
        sound,
        playing,
        setPlaying,
        volume,
        setVolume,
        muted,
        setMuted,
        handleLike,
        loadingAudio,
        setLoadingAudio,
      }}
    >
      {children}
    </AudioContext.Provider>
  );
}

export function useAudio() {
  const context = useContext(AudioContext);

  if (context === undefined) {
    throw new Error("useAudio must be used within a AudioProvider");
  }

  return context;
}
