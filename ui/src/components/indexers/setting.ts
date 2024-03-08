export interface Setting {
  setting: string;
  value: boolean;
}

export interface SettingsArgs {
  id: string;
  setting: Setting;
}
