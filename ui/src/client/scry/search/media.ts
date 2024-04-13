import { SearchResponse } from './search';

/*
type Media struct {
	ID          string    `json:"id"`
	Type        string    `json:"type"`
	Name        string    `json:"name,omitempty"`
	Kind        string    `json:"kind"`
	Source      string    `json:"source,omitempty"`
	SourceID    string    `json:"source_id,omitempty"`
	SearchName  string    `json:"search_name,omitempty"`
	Display     string    `json:"display,omitempty"`
	Title       string    `json:"title,omitempty"`
	Description string    `json:"description,omitempty"`
	Link        string    `json:"link,omitempty"`
	ReleaseDate string    `json:"release_date,omitempty"`
	Background  string    `json:"background,omitempty"`
	Cover       string    `json:"cover,omitempty"`
	Created     time.Time `json:"created_at,omitempty"`
	Updated     time.Time `json:"updated_at,omitempty"`
}
*/
export interface Media {
  id: string;
  type: string;
  name: string;
  kind: string;
  source: string;
  sourceid: string;
  searchname: string;
  display: string;
  title: string;
  description: string;
  link: string;
  releasedate: string;
  background: string;
  cover: string;
  created_at: string;
  updated_at: string;
}
export interface MediaSearchResponse extends SearchResponse {
  Media: Media[];
}
