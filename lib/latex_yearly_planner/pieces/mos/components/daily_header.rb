# frozen_string_literal: true

module LatexYearlyPlanner
  module Pieces
    module Mos
      module Components
        class DailyHeader < Header
          def generate(day)
            "#{top_table}" \
              '\\hfill{}' \
              "#{title(day)}" \
              '\\hrule{}' \
              "#{margin_note(highlight_quarters: [day.quarter], highlight_months: [day.month])}" \
              "\\medskip\n"
          end

          def title(day)
            TeX::TextSize.new(day.name).huge
          end
        end
      end
    end
  end
end