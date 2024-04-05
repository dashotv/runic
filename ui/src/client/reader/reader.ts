// TODO: implement interfaces/types for reader
export interface Source {
  Name: string;
  URL: string;
  Key: string;
  UserID: string;
  Insecure: boolean;
  Type: string;
  Caps: SourceCaps;
}

export interface SourceCaps {
  Server: SourceCapsServer;
  Limits: SourceCapsLimits;
  Searching: SourceCapsSearching;
  Categories: SourceCapsCategories;
}
export interface SourceCapsServer {
  Title: string;
}
export interface SourceCapsLimits {
  Default: string;
  Max: string;
}
export interface SourceCapsSearching {
  Search: SourceCapsSearchingSearch;
  Tvsearch: SourceCapsSearchingSearch;
  Moviesearch: SourceCapsSearchingSearch;
  Musicsearch: SourceCapsSearchingSearch;
  Audiosearch: SourceCapsSearchingSearch;
  Booksearch: SourceCapsSearchingSearch;
}
export interface SourceCapsSearchingSearch {
  Available: string;
  SupportedParams: string[];
  SearchEngine: string;
}
export interface SourceCapsCategories {
  Category: SourceCapsCategoriesCategory[];
}
export interface SourceCapsCategoriesCategory {
  ID: string;
  Name: string;
  Subcat: SourceCapsCategoriesCategorySubcat[];
}
export interface SourceCapsCategoriesCategorySubcat {
  ID: string;
  Name: string;
}
