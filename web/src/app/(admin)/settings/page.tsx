"use client";

import { PageTitle } from "../page-title";
import { Switch } from "@/components/ui/switch";
import { Button } from "@/components/ui/button";
import {
  useGetConfigQuery,
  useUpdateConfigMutation,
} from "@/services/config-service";
import { useState, useEffect } from "react";
import { Config } from "@/types/config";

export default function SettingsPage() {
  const { data: config } = useGetConfigQuery();
  const [updateConfig, { isLoading }] = useUpdateConfigMutation();

  const [settings, setSettings] = useState<Config>({
    id: 1,
    email_domain: false,
    validate_phone: false,
    status_rules: false,
    delete_limit: false,
  });

  useEffect(() => {
    if (config) {
      setSettings(config.data);
    }
  }, [config]);

  const handleChange = (key: string, value: number | boolean) => {
    setSettings((prev) => ({ ...prev, [key]: value }));
  };

  const handleSave = async () => {
    await updateConfig(settings);
  };

  return (
    <>
      <PageTitle title="Settings" />
      <div className="p-6 space-y-4 bg-white shadow rounded-lg">
        <div className="flex justify-between items-center">
          <span>Allowed email domain</span>
          <Switch
            checked={settings.email_domain}
            onCheckedChange={(checked) => handleChange("email_domain", checked)}
          />
        </div>

        <div className="flex justify-between items-center">
          <span>Validate phone number</span>
          <Switch
            checked={settings.validate_phone}
            onCheckedChange={(checked) =>
              handleChange("validate_phone", checked)
            }
          />
        </div>

        <div className="flex justify-between items-center">
          <span>Apply student status change rules</span>
          <Switch
            checked={settings.status_rules}
            onCheckedChange={(checked) => handleChange("status_rules", checked)}
          />
        </div>

        <div className="flex justify-between items-center">
          <span>Student deletion limit (minutes)</span>
          <Switch
            checked={settings.delete_limit}
            onCheckedChange={(checked) => handleChange("delete_limit", checked)}
          />
        </div>

        <div className="flex justify-end">
          <Button onClick={handleSave} disabled={isLoading || config?.data === settings}>
            Save
          </Button>
        </div>
      </div>
    </>
  );
}
