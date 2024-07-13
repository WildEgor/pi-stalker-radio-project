export type Station = {
  name: string;
  serveruuid: string;
  country: string;
  language: string;
  url: string;
  votes: number;
  favicon: string;
  tags: string;
  stationuuid: string;
};

export type StationLocation = {
  name: string;
  serveruuid: string;
  tags: string;
  url: string;
  language: string;
  votes: number;
  favicon: string;
  country: string;
  geo_lat: number;
  geo_long: number;
};
