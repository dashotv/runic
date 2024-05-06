// Code generated by github.com/dashotv/golem. DO NOT EDIT.
import { Batch, BatchResult, Response, runicClient } from '.';
import * as parser from './parser';

export interface ParserParseRequest {
  title: string;
  type: string;
}
export interface ParserParseResponse extends Response {
  result: parser.TorrentInfo;
}
export const ParserParse = async (params: ParserParseRequest) => {
  const response = await runicClient.get(`/parser/parse?title=${params.title}&type=${params.type}`);

  if (!response.data) {
    throw new Error('response empty?');
  }

  if (response.data.error) {
    if (response.data.message) {
      throw new Error(response.data.message);
    }
    throw new Error('unknown error');
  }

  return response.data as ParserParseResponse;
};

export interface ParserTitleRequest {
  title: string;
  type: string;
}
export interface ParserTitleResponse extends Response {
  result: parser.TorrentInfo;
}
export const ParserTitle = async (params: ParserTitleRequest) => {
  const response = await runicClient.get(`/parser/title?title=${params.title}&type=${params.type}`);

  if (!response.data) {
    throw new Error('response empty?');
  }

  if (response.data.error) {
    if (response.data.message) {
      throw new Error(response.data.message);
    }
    throw new Error('unknown error');
  }

  return response.data as ParserTitleResponse;
};

export interface ParserBatchRequest {
  batch: Batch;
}
export interface ParserBatchResponse extends Response {
  result: BatchResult[];
}
export const ParserBatch = async (params: ParserBatchRequest) => {
  const response = await runicClient.post(`/parser/batch?`, params.batch);

  if (!response.data) {
    throw new Error('response empty?');
  }

  if (response.data.error) {
    if (response.data.message) {
      throw new Error(response.data.message);
    }
    throw new Error('unknown error');
  }

  return response.data as ParserBatchResponse;
};
