import {useAudio} from "@/store/providers/audioContext.tsx";
import {Popover, PopoverContent, PopoverTrigger} from "@/atoms/popover.tsx";
import {Button} from "@/atoms/button.tsx";
import {Volume2, VolumeX} from "lucide-react";
import {Slider} from "@/atoms/slider.tsx";

const VolumeControl = () => {
  const { sound, setMuted, muted } = useAudio();

  const handleMute = () => {
    sound?.mute(!sound?.mute());
    setMuted(sound?.mute() ? true : false);
  };
  return (
    <Popover>
      <PopoverTrigger>
        <Button
          variant={"unstyled"}
          className="p-2 text-foreground hover:text-accent"
        >
          <Volume2 />
        </Button>
      </PopoverTrigger>
      <PopoverContent className="flex gap-3 z-30">
        <Button
          onClick={handleMute}
          variant="ghost"
          className={`p-2 ${muted ? "bg-primary" : " text-foreground/50"}`}
        >
          <VolumeX />
        </Button>
        <Slider
          onValueChange={(value) => sound?.volume(value[0])}
          orientation="horizontal"
          defaultValue={[0.5]}
          max={1}
          step={0.05}
          className=""
        />
      </PopoverContent>
    </Popover>
  );
};

export default VolumeControl;
