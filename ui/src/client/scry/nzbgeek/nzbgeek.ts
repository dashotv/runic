// TODO: implement interfaces/types for nzbgeek

export interface SearchResult {
  title: string;
  guid: string;
  link: string;
  comments: string;
  pubDate: string;
  category: string;
  description: string;
  enclosure: {
    '@attributes': {
      url: string;
      length: string;
      type: string;
    };
  };
  attributes: {
    attribute: {
      name: string;
      value: string;
    };
  }[];
}

export interface NzbgeekFormTv {
  tvdbid?: string;
  season?: number | '';
  episode?: number | '';
}
export interface NzbgeekFormMovie {
  imdbid?: string;
  tmdbid?: string;
}
export interface NzbResponse {
  Version: string;
  Error: string;
  Timestamp: string;
  Status: NzbResponseStatus;
  Result: Nzb[];
}

export interface Nzb {
  nzbid: number;
  RemainingSizeMB: number; // 3497
  PausedSizeMB: number; // 3497
  RemainingFileCount: number; // 73
  RemainingParCount: number; // 9
  MinPriority: number; // 0
  MaxPriority: number; // 0
  ActiveDownloads: number; // 0
  Status: string; // PAUSED
  NZBName: string; // Brave.Are.the.Fallen.2020.1080p.AMZN.WEB-DL.DDP2.0.H.264-ExREN,
  NZBNicename: string; // Brave.Are.the.Fallen.2020.1080p.AMZN.WEB-DL.DDP2.0.H.264-ExREN,
  Kind: string; // NZB
  URL: string; // ,
  NZBFilename: string; // Brave.Are.the.Fallen.2020.1080p.AMZN.WEB-DL.DDP2.0.H.264-ExREN,
  DestDir: string; // /data/intermediate/Brave.Are.the.Fallen.2020.1080p.AMZN.WEB-DL.DDP2.0.H.264-ExREN.#4,
  FinalDir: string; // ,
  Category: string; // ,
  ParStatus: string; // NONE
  ExParStatus: string; // NONE
  UnpackStatus: string; // NONE
  MoveStatus: string; // NONE
  ScriptStatus: string; // NONE
  DeleteStatus: string; // NONE
  MarkStatus: string; // NONE
  UrlStatus: string; // NONE
  FileSizeMB: number; // 3651
  FileCount: number; // 77
  MinPostTime: number; // 1586073677
  MaxPostTime: number; // 1586073793
  TotalArticles: number; // 4992
  SuccessArticles: number; // 212
  FailedArticles: number; // 0
  Health: number; // 1000
  CriticalHealth: number; // 898
  DupeKey: string; //
  DupeScore: number; // 0
  DupeMode: string; // SCORE
  Deleted: boolean;
  DownloadedSizeMB: number; // 235
  DownloadTimeSec: number; // 44
  PostTotalTimeSec: number; // 0
  ParTimeSec: number; // 0
  RepairTimeSec: number; // 0
  UnpackTimeSec: number; // 0
  MessageCount: number; // 95
  ExtraParBlocks: number; // 0

  // Parameters         []Parameter
  // ScriptStatuses     []ScriptStatus
  // ServerStats        []ServerStat
  PostInfoText: string; // NONE
  PostStageProgress: number; // 9193728
  PostStageTimeSec: number; // 0
}
export interface NzbResponseStatus {
  DownloadRate: number;
  FreeDiskSpaceMB: number;
}
