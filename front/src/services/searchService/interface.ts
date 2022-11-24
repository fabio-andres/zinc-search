interface RootObject {
    took: number;
    timed_out: boolean;
    _shards: Shards;
    hits: Hits;
  }
  
  interface Hits {
    total: Total;
    max_score: number;
    hits: Hit[];
  }
  
  interface Hit {
    _index: string;
    _type: string;
    _id: string;
    _score: number;
    '@timestamp': string;
    _source: Source;
  }
  
  interface Source {
    Body: string;
    'Content-Transfer-Encoding': string;
    'Content-Type': string;
    Date: string;
    From: string;
    'Message-ID': string;
    'Mime-Version': string;
    Subject: string;
    To: string;
    'X-FileName': string;
    'X-Folder': string;
    'X-From': string;
    'X-Origin': string;
    'X-To': string;
    'X-bcc': string;
    'X-cc': string;
  }
  
  interface Total {
    value: number;
  }
  
  interface Shards {
    total: number;
    successful: number;
    skipped: number;
    failed: number;
  }

  export default RootObject